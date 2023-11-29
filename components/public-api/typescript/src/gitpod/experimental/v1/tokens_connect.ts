/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file gitpod/experimental/v1/tokens.proto (package gitpod.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreatePersonalAccessTokenRequest, CreatePersonalAccessTokenResponse, DeletePersonalAccessTokenRequest, DeletePersonalAccessTokenResponse, GetPersonalAccessTokenRequest, GetPersonalAccessTokenResponse, ListPersonalAccessTokensRequest, ListPersonalAccessTokensResponse, RegeneratePersonalAccessTokenRequest, RegeneratePersonalAccessTokenResponse, UpdatePersonalAccessTokenRequest, UpdatePersonalAccessTokenResponse } from "./tokens_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service gitpod.experimental.v1.TokensService
 */
export const TokensService = {
  typeName: "gitpod.experimental.v1.TokensService",
  methods: {
    /**
     * CreatePersonalAccessTokenRequest creates a new token.
     *
     * @generated from rpc gitpod.experimental.v1.TokensService.CreatePersonalAccessToken
     */
    createPersonalAccessToken: {
      name: "CreatePersonalAccessToken",
      I: CreatePersonalAccessTokenRequest,
      O: CreatePersonalAccessTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListPersonalAccessTokens returns token by ID.
     *
     * @generated from rpc gitpod.experimental.v1.TokensService.GetPersonalAccessToken
     */
    getPersonalAccessToken: {
      name: "GetPersonalAccessToken",
      I: GetPersonalAccessTokenRequest,
      O: GetPersonalAccessTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListPersonalAccessTokens returns a list of tokens.
     *
     * @generated from rpc gitpod.experimental.v1.TokensService.ListPersonalAccessTokens
     */
    listPersonalAccessTokens: {
      name: "ListPersonalAccessTokens",
      I: ListPersonalAccessTokensRequest,
      O: ListPersonalAccessTokensResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RegeneratePersonalAccessToken generates a new token and replaces the previous one.
     *
     * @generated from rpc gitpod.experimental.v1.TokensService.RegeneratePersonalAccessToken
     */
    regeneratePersonalAccessToken: {
      name: "RegeneratePersonalAccessToken",
      I: RegeneratePersonalAccessTokenRequest,
      O: RegeneratePersonalAccessTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdatePersonalAccessToken updates writable properties of a PersonalAccessToken.
     *
     * @generated from rpc gitpod.experimental.v1.TokensService.UpdatePersonalAccessToken
     */
    updatePersonalAccessToken: {
      name: "UpdatePersonalAccessToken",
      I: UpdatePersonalAccessTokenRequest,
      O: UpdatePersonalAccessTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeletePersonalAccessToken removes token by ID.
     *
     * @generated from rpc gitpod.experimental.v1.TokensService.DeletePersonalAccessToken
     */
    deletePersonalAccessToken: {
      name: "DeletePersonalAccessToken",
      I: DeletePersonalAccessTokenRequest,
      O: DeletePersonalAccessTokenResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
