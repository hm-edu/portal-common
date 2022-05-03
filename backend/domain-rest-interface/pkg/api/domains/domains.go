package domains

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/hm-edu/domain-rest-interface/ent"
	"github.com/hm-edu/domain-rest-interface/ent/delegation"
	"github.com/hm-edu/domain-rest-interface/pkg/model"
	"github.com/hm-edu/portal-common/auth"
	"github.com/hm-edu/portal-common/helper"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// ListDomains godoc
// @Summary List domains.
// @Description Lists all domains that are either owned or delegated, or a child of a owned or delegated domain.
// @Tags Domains
// @Accept json
// @Produce json
// @Router /domains/ [get]
// @Security API
// @Success 200 {object} []model.Domain
// @Response default {object} echo.HTTPError "Error processing the request"
func (h *Handler) ListDomains(c echo.Context) error {
	domains, err := h.domainStore.ListDomains(c.Request().Context(), auth.UserFromRequest(c), false)
	if err != nil {
		h.logger.Error("Listing domains failed", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err, Message: "Invalid Request"}
	}
	return c.JSON(http.StatusOK, helper.Map(domains, model.DomainToOutput))
}

// CreateDomain godoc
// @Summary Create a domain.
// @Description Creates a new domain if the FQDN is not already taken. Approvement is automatically done, in case the user owns a upper zone or a upper zone was already delegated to him.
// @Tags Domains
// @Accept json
// @Produce json
// @Router /domains/ [post]
// @Param domain body model.DomainRequest true "The Domain to create"
// @Security API
// @Success 201 {object} model.Domain
// @Response default {object} echo.HTTPError "Error processing the request"
func (h *Handler) CreateDomain(c echo.Context) error {
	req := &model.DomainRequest{}
	if err := req.Bind(c, h.validator); err != nil {
		h.logger.Error("Binding request failed", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err, Message: "Invalid Request"}
	}
	domain := ent.Domain{Owner: auth.UserFromRequest(c), Fqdn: req.FQDN}

	domains, err := h.domainStore.ListDomains(c.Request().Context(), auth.UserFromRequest(c), true)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err, Message: "Invalid Request"}
	}

	if helper.Any(domains, func(i *ent.Domain) bool { return i.Approved && strings.HasSuffix(domain.Fqdn, "."+i.Fqdn) }) {
		domain.Approved = true
	}

	created, err := h.domainStore.Create(c.Request().Context(), &domain)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest}
	}
	return c.JSON(http.StatusCreated, model.DomainToOutput(created))
}

// DeleteDomains godoc
// @Summary Delete a domain
// @Description Deletes a domain. Existing certificates are not are not longer accessible.
// @Tags Domains
// @Accept json
// @Produce json
// @Router /domains/{id} [delete]
// @Param id path int true "Domain ID"
// @Security API
// @Success 204
// @Response default {object} echo.HTTPError "Error processing the request"
func (h *Handler) DeleteDomains(c echo.Context) error {

	_, domain, err := h.extractData(c)
	if err != nil {
		return err
	}

	if err := h.domainStore.Delete(c.Request().Context(), []*ent.Domain{domain}); err != nil {
		h.logger.Error("Deleting domain failed", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Error while deleting domains"}
	}
	return c.NoContent(http.StatusNoContent)
}

// ApproveDomain godoc
// @Summary Approve domain request
// @Description Approves an outstanding domain request
// @Tags Domains
// @Accept json
// @Produce json
// @Router /domains/{id}/approve [post]
// @Param id path int true "Domain ID"
// @Security API
// @Success 200 {object} model.Domain The updated domain
// @Response default {object} echo.HTTPError "Error processing the request"
func (h *Handler) ApproveDomain(c echo.Context) error {

	_, domain, err := h.extractData(c)
	if err != nil {
		return err
	}

	updated, err := h.domainStore.Approve(c.Request().Context(), domain)
	if err != nil {
		h.logger.Error("Approving domain failed", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Error while approving domain"}
	}

	return c.JSON(http.StatusOK, model.DomainToOutput(updated))
}

func (h *Handler) extractData(c echo.Context) ([]*ent.Domain, *ent.Domain, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error("Invalid domain id", zap.Error(err))
		return nil, nil, &echo.HTTPError{Code: http.StatusBadRequest, Message: "Invalid domain ID"}
	}

	domains, err := h.domainStore.ListDomains(c.Request().Context(), auth.UserFromRequest(c), true)
	if err != nil {
		h.logger.Error("Listing domains failed", zap.Error(err))
		return nil, nil, &echo.HTTPError{Code: http.StatusInternalServerError, Message: "Error while fetching domains"}
	}

	var domain *ent.Domain
	domain = helper.First(domains, func(i *ent.Domain) bool { return i.ID == id })
	if domain == nil {
		return nil, nil, &echo.HTTPError{Code: http.StatusNotFound, Message: "Domain not found"}
	}
	return domains, domain, nil
}

// TransferDomain godoc
// @Summary Transfer domain
// @Description Transfers a domain to a new owner. Transferring is only possible if you are either the owner of the domain itself or responsible for one of the parent zones.
// @Tags Domains
// @Accept json
// @Produce json
// @Router /domains/{id}/transfer [post]
// @Param domain body model.TransferRequest true "The Domain to transfer"
// @Param id path int true "Domain ID"
// @Security API
// @Success 200 {object} model.Domain The updated domain
// @Response default {object} echo.HTTPError "Error processing the request"
func (h *Handler) TransferDomain(c echo.Context) error {

	req := &model.TransferRequest{}
	if err := req.Bind(c, h.validator); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest}
	}

	domains, domain, err := h.extractData(c)
	if err != nil {
		return err
	}

	if !helper.Any(domains, func(i *ent.Domain) bool {
		return i.Approved && (strings.HasSuffix(domain.Fqdn, "."+i.Fqdn) || (domain.Fqdn == i.Fqdn && req.Owner == auth.UserFromRequest(c)))
	}) {
		return &echo.HTTPError{Code: http.StatusForbidden, Message: "You are not responsible for this zone"}
	}

	updated, err := h.domainStore.Owner(c.Request().Context(), domain, req.Owner)
	if err != nil {
		h.logger.Error("Transferring domain failed", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusBadRequest}
	}

	return c.JSON(http.StatusOK, model.DomainToOutput(updated))
}

// DeleteDelegation godoc
// @Summary Delete delegation.
// @Description Deletes an existing delegation.
// @Tags Domains
// @Accept json
// @Produce json
// @Router /domains/{id}/delegation/{delegation} [delete]
// @Param id path int true "Domain ID"
// @Param delegation path int true "Delegation ID"
// @Security API
// @Success 200 {object} model.Domain The updated domain
// @Response default {object} echo.HTTPError "Error processing the request"
func (h *Handler) DeleteDelegation(c echo.Context) error {
	_, domain, err := h.extractData(c)
	if err != nil {
		return err
	}

	delegationID, err := strconv.Atoi(c.Param("delegation"))
	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest}
	}

	delegated, err := domain.QueryDelegations().Where(delegation.ID(delegationID)).First(c.Request().Context())
	if err != nil {
		h.logger.Error("Delegation not found", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Delegation not found"}
	}

	updated, err := h.domainStore.DeleteDelegation(c.Request().Context(), domain, delegated)
	if err != nil {
		h.logger.Error("Deleting delegation failed", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Deleting delegation failed"}
	}

	return c.JSON(http.StatusOK, model.DomainToOutput(updated))
}

// AddDelegation godoc
// @Summary Add delegation.
// @Description Adds a new delegation to an existing domain.
// @Tags Domains
// @Accept json
// @Produce json
// @Router /domains/{id}/delegation [post]
// @Param delegation body model.DelegationRequest true "The Delegation to add"
// @Param id path int true "Domain ID"
// @Security API
// @Success 200 {object} model.Domain The updated domain
// @Response default {object} echo.HTTPError "Error processing the request"
func (h *Handler) AddDelegation(c echo.Context) error {

	req := &model.DelegationRequest{}
	if err := req.Bind(c, h.validator); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Internal: err, Message: "Invalid Request"}
	}

	_, domain, err := h.extractData(c)
	if err != nil {
		return err
	}

	updated, err := h.domainStore.AddDelegation(c.Request().Context(), domain, req.User)
	if err != nil {
		h.logger.Error("Adding delegation failed", zap.Error(err))
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "Adding delegation failed"}
	}

	return c.JSON(http.StatusOK, model.DomainToOutput(updated))
}
