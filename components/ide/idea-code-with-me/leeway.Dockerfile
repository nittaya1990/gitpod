# Copyright (c) 2021 Gitpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License-AGPL.txt in the project root for license information.

FROM golang:1.17 AS build
WORKDIR /app
COPY status/* /app/
RUN go build -o status

FROM scratch
COPY --chown=33333:33333 startup.sh supervisor-ide-config.json /ide-desktop/
COPY --chown=33333:33333 --from=build /app/status /ide-desktop/
