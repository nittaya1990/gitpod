#!/bin/bash -li
# Copyright (c) 2021 Gitpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License-AGPL.txt in the project root for license information.

/ide-desktop/status 24000 &

export IDEA_VERSION=213.4101
url=$(echo "$IDEA_IU_URL" | sed "s/213.2899/$IDEA_VERSION/g")

mkdir -p /home/gitpod/.local/share/JetBrains/CwmHost/${IDEA_VERSION}/ && \
    cd /home/gitpod/.local/share/JetBrains/CwmHost/${IDEA_VERSION}/ && \
    wget --quiet -O ideaIU-${IDEA_VERSION}.tar.gz "$url" && \
    tar xf ideaIU-${IDEA_VERSION}.tar.gz && \
    rm -rf ideaIU-${IDEA_VERSION}.tar.gz

export IDEA=/home/gitpod/.local/share/JetBrains/CwmHost/${IDEA_VERSION}/idea-IU-${IDEA_VERSION}/
export CWM_NON_INTERACTIVE=1
export CWM_HOST_PASSWORD=gitpod
export CWM_HOST_STATUS_OVER_HTTP_TOKEN=gitpod
${IDEA}/bin/remote-dev-server.sh cwmHost "$GITPOD_REPO_ROOT" > /home/gitpod/jetbrains-logs.txt 2>&1
