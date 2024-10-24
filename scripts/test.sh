mods=$(go list -f '{{.Dir}}' -m)
for mod in $mods; do
            go test  -coverprofile $mod.out "$mod" ./...
done