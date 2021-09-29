// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the Gitpod Enterprise Source Code License,
// See License.enterprise.txt in the project root folder.

package agent_test

import (
	"bytes"
	"testing"

	"github.com/gitpod-io/gitpod/agent-smith/pkg/agent"
)

func TestParsePsOutput(t *testing.T) {
	output := `      1       0 supervisor      supervisor run
   36       1 gitpod-node     /ide/node/bin/gitpod-node ./out/gitpod.js /workspace/gitpod/gitpod-ws.code-workspace --port 23000 --hostname 0.0.0.0
   47       1 dropbear        /.supervisor/dropbear/dropbear -F -E -w -s -p :23001 -r /tmp/hostkey3308595158
  269       1 Xvfb            Xvfb -screen 0 1920x1080x16 -ac -pn -noreset
  270       1 openbox         openbox
  311       1 x11vnc          x11vnc -localhost -shared -display :0 -forever -rfbport 5900 -bg -o /tmp/x11vnc-0.log
  312       1 start-vnc-sessi /bin/bash /usr/bin/start-vnc-session.sh
  314     312 bash            bash ./novnc_proxy --vnc localhost:5900 --listen 6080
  323     314 python3         /home/gitpod/.pyenv/versions/3.8.11/bin/python3 -m websockify --web /opt/novnc/utils/../ 6080 localhost:5900
  700       1 bash            /bin/bash
  701       1 bash            /bin/bash
  704       1 bash            /bin/bash
 1572     700 leeway          leeway run components/ee/agent-smith:qemu
 1761    1572 bash            bash /tmp/459992179.sh
 1762    1761 bash            bash scripts/qemu.sh
 1770    1762 sudo            sudo qemu-system-x86_64 -kernel /boot/vmlinuz-5.4.0-1046-gke -boot c -m 2049M -hda /root/_output/bionic-server-cloudimg-amd64.qcow2 -net user -smp 2 -append root=/dev/sda rw console=ttyS0,115200 acpi=off nokaslr -nic user,hostfwd=tcp::2222-:22 -s -serial mon:stdio -display none
 1771    1770 qemu-system-x86 qemu-system-x86_64 -kernel /boot/vmlinuz-5.4.0-1046-gke -boot c -m 2049M -hda /root/_output/bionic-server-cloudimg-amd64.qcow2 -net user -smp 2 -append root=/dev/sda rw console=ttyS0,115200 acpi=off nokaslr -nic user,hostfwd=tcp::2222-:22 -s -serial mon:stdio -display none
28796      36 gitpod-node     /ide/node/bin/gitpod-node /ide/out/bootstrap-fork --type=watcherService
28807      36 gitpod-node     /ide/node/bin/gitpod-node /ide/out/bootstrap-fork --type=extensionHost --uriTransformerPath=/ide/out/serverUriTransformer
28896   28807 gopls           /home/gitpod/go-packages/bin/gopls -mode=stdio
28910   28807 gitpod-node     /ide/node/bin/gitpod-node /ide/extensions/redhat.vscode-yaml/node_modules/yaml-language-server/out/server/src/server.js --node-ipc --clientProcessId=28807
29152   28807 gitpod-node     /ide/node/bin/gitpod-node /ide/extensions/json-language-features/server/dist/node/jsonServerMain --node-ipc --clientProcessId=28807
29954      36 gitpod-node     /ide/node/bin/gitpod-node /ide/out/bootstrap-fork --type=watcherService
29965      36 gitpod-node     /ide/node/bin/gitpod-node /ide/out/bootstrap-fork --type=extensionHost --uriTransformerPath=/ide/out/serverUriTransformer
30031   29965 gopls           /home/gitpod/go-packages/bin/gopls -mode=stdio
31193   29965 gitpod-node     /ide/node/bin/gitpod-node /home/gitpod/.gitpod-code/extensions/matthewpi.caddyfile-support-0.2.0/server/out/server.js --node-ipc --clientProcessId=29965
31484   29965 gitpod-node     /ide/node/bin/gitpod-node /ide/extensions/redhat.vscode-yaml/node_modules/yaml-language-server/out/server/src/server.js --node-ipc --clientProcessId=29965
33105   29965 gitpod-node     /ide/node/bin/gitpod-node /ide/extensions/json-language-features/server/dist/node/jsonServerMain --node-ipc --clientProcessId=29965
33681   29965 gitpod-node     /ide/node/bin/gitpod-node /home/gitpod/.gitpod-code/extensions/bradlc.vscode-tailwindcss-0.6.14/dist/server/tailwindServer.js --node-ipc --clientProcessId=29965
101911       1 go <defunct>    [go] <defunct>
102484       1 go <defunct>    [go] <defunct>
117495       1 go <defunct>    [go] <defunct>
126718     704 ps              ps -e -o pid,ppid,comm,args --no-headers
`

	m, err := agent.ParsePsOutput(bytes.NewBufferString(output))
	if err != nil {
		t.Fatal(err)
	}

	p1 := m.GetByPID(1)
	if p1 == nil {
		t.Fatal("expected process 1 to be present")
	}
	p0, ok := m.GetParent(p1.PID)
	if !ok {
		t.Fatal("expected process 1 to still be present")
	}
	if p0 != nil {
		t.Fatal("expected process 0 not to be present")
	}

	p28910 := m.GetByPID(28910)
	if p28910 == nil {
		t.Fatal("expected process 28910 to be present")
	}
	if p28910.PPID != 28807 {
		t.Fatal("expected process 28910 parent to be 28807")
	}
}
