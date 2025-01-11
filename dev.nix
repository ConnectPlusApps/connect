{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gopls
    pkgs.delve
    pkgs.nodejs
    pkgs.yarn
    (pkgs.flutter.mkFlutter {
      version = "3.16.0";
      channel = "stable";
    })
    pkgs.postgresql_14
    pkgs.xdg-utils
  ];
}
