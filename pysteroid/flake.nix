{
  description = "pysteroid dev flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }: {
    devShell.x86_64-linux = let
      pkgs = nixpkgs.legacyPackages.x86_64-linux;
    in pkgs.mkShell {
      packages = with pkgs; [ uv (python313.withPackages (ps: [ ps.pygame ])) ];
    };
  };
}
