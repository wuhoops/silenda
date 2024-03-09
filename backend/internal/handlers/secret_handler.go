package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"github.com/wuhoops/silenda/backend/internal/services"
)

type SecretHandler struct {
	secretService services.SecretService
}

func NewSecretHandler(secretService services.SecretService) SecretHandler {
	return SecretHandler{
		secretService: secretService,
	}
}

// FindSecretsByEnvMode godoc
// @tags secret
// @id findSecretsByEnvMode
// @summary Find secrets by env mode
// @accept json
// @produce json
// @Param payload body FindSecretsByEnvModeBody true "FindSecretsByEnvModeBody"
// @Success 200 {object} Response[[]FindSecretsByEnvModeResponse]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.findSecretsByEnvMode [post]
func (s *SecretHandler) FindSecretsByEnvMode(c *fiber.Ctx) error {
	var body FindSecretsByEnvModeBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	workspaceIdUuid, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	secrets, err := s.secretService.FindSecretsByEnvMode(workspaceIdUuid, body.EnvMode)
	if err != nil {
		return err
	}
	var response []FindSecretsByEnvModeResponse
	for _, secret := range secrets {
		response = append(response, FindSecretsByEnvModeResponse{
			ID:    (*secret.ID).String(),
			Key:   *secret.Key,
			Value: *secret.Value,
		})
	}

	return Ok(c, response)
}

// CreateSecret godoc
// @tags secret
// @id createSecret
// @summary Create secret
// @accept json
// @produce json
// @Param payload body CreateSecretBody true "CreateSecretBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.createSecret [post]
func (s *SecretHandler) CreateSecret(c *fiber.Ctx) error {
	var body CreateSecretBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	UuidWorkspaceIdUuid, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	secret := domains.Secret{
		WorkspaceID: &UuidWorkspaceIdUuid,
		EnvMode:     (*domains.EnvMode)(&body.EnvMode),
		Key:         &body.Key,
		Value:       &body.Value,
	}
	if err := s.secretService.CreateSecret(secret); err != nil {
		return err
	}
	return Created(c, "Create secret successfully!")
}

// DeleteSecretById godoc
// @tags secret
// @id deleteSecretById
// @summary Delete secret by id
// @accept json
// @produce json
// @Param payload body DeleteSecretByIdBody true "DeleteSecretByIdBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.deleteSecretByID [post]
func (s *SecretHandler) DeleteSecretById(c *fiber.Ctx) error {
	var body DeleteSecretByIdBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	secretIdUuid, err := uuid.Parse(body.SecretID)
	if err != nil {
		return err
	}
	err = s.secretService.DeleteSecretById(secretIdUuid)
	if err != nil {
		return err
	}
	return Ok(c, "delete secret successfully")
}
