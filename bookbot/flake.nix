{
  description = "boot.dev bookbot — Python";

  inputs.nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";

  outputs =
    { self, nixpkgs }:
    let
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
    in
    {
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            python313
            uv
            pyright
            ruff
          ];

          shellHook = ''
            export UV_PYTHON="$(which python3)"
            export UV_PYTHON_PREFERENCE="only-system"
          '';
        };
      });
    };
}
