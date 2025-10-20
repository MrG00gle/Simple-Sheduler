{
  pkgs,
  lib,
  config,
  ...
}:
{
  # https://devenv.sh/languages/
  languages = {
    go.enable = true;
    c.enable = true;
  };

  # https://devenv.sh/packages/
  packages = [
    pkgs.go-bindata
    pkgs.gopy
    pkgs.gcc
  ];

  # Additional setup for CGO
  env.CGO_ENABLED = "1";

  # See full reference at https://devenv.sh/reference/options/
}
