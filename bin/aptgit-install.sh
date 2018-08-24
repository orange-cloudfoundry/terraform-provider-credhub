  

grep -Fxq "providers {" ~/.terraformrc &> /dev/null
if [[ $? != 0 ]]; then
	cat <<EOF >> ~/.terraformrc
providers {
{{.Metadata.shortname}} = "{{.Home}}/.terraform.d/providers/{{.Name}}"
}
EOF
else
	grep -Fxq "{{.Metadata.shortname}}" ~/.terraformrc &> /dev/null
	if [[ $? != 0 ]]; then
		echo "{{.Name}}-{{.Version}} has been installed."
		exit 0
	fi
	awk '/providers {/ { print; print "{{.Metadata.shortname}} = \"provider_path\""; next }1' ~/.terraformrc > /tmp/.terraformrc
	mv /tmp/.terraformrc ~/
fi
echo "{{.Name}}-{{.Version}} has been installed."
