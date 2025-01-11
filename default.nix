{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.flutter37
    # Add other dependencies here
  ];
}
