{
  description = "godex flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        go = pkgs.go;
      in
      {
        packages.default = pkgs.buildGoModule rec {
          pname = "godex";
          version = "0.1.0";
          inherit (pkgs.lib) stdenv;

          # Use the repository root as source
          src = ./.;

          # If you have go.mod/go.sum, buildGoModule will use them automatically.
          goPackagePath = "example.com/example-go-tool";

          vendorSha256 = null;

          meta = with pkgs.lib; {
            description = "go pokedex REPL";
            license = licenses.mit;
            maintainers = with maintainers; [];
          };
        };

        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            golangci-lint
            delve
            goreleaser
          ];
        };
      });
}

