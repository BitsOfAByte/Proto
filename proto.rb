# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Proto < Formula
  desc "Proto compatability tool manager
"
  homepage "https://github.com/BitsOfAByte/proto"
  version "0.12.2"
  license "GPL-3.0-only"

  depends_on "gnu-tar"
  depends_on :linux

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.12.2/proto_0.12.2_linux_armv6.zip"
      sha256 "8ba9f3145ef94919f211e3a1081fba0da72d2aeef1e0aac248e7e6de14e25c42"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.12.2/proto_0.12.2_linux_amd64.zip"
      sha256 "50f2b890649f19d931a2a6c8bc3421c19cf7cd886c4e30601f58f9f6f97bf859"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v0.12.2/proto_0.12.2_linux_arm64.zip"
      sha256 "07e3048f3a6f156b08b40fa3e3fffe3f4ab86f868e9c876ffee68ea0841fa171"

      def install
        bin.install "proto"
      end
    end
  end
end
