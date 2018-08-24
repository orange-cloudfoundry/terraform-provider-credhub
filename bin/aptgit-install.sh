#!/usr/bin/env bash 
grep -Fxq "providers {" {{.Home}}/.terraformrc &> /dev/null
if [[ $? != 0 ]]; then
cat <<EOF >> ~/.terraformrc
providers {
{{.Metadata.shortname}} = "{{.Home}}/.terraform.d/providers/{{.Name}}"
}
EOF
else
	grep -Fxq "{{.Metadata.shortname}}" {{.Home}}/.terraformrc &> /dev/null
	if [[ $? == 0 ]]; then
	  echo "{{.Name}}-{{.Version}} has been installed."
	  exit 0
	fi
	awk '/providers {/ { print; print "{{.Metadata.shortname}} = \"{{.InstallPath}}\""; next }1' {{.Home}}/.terraformrc > /tmp/.terraformrc
	mv /tmp/.terraformrc {{.Home}}/.terraformrc
fi
echo "{{.Name}}-{{.Version}} has been installed and added to list."