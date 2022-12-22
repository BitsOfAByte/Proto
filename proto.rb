# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Proto < Formula
  desc "Proto compatability tool manager
"
  homepage "https://github.com/BitsOfAByte/proto"
  version "1.1.0"
  license "GPL-3.0-only"

  depends_on "gnu-tar"
  depends_on :linux

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/BitsOfAByte/proto/releases/download/v1.1.0/proto_1.1.0_linux_amd64.zip"
      sha256 "ee58033cc4ed2d5e62120b4a8d86166fd00bc6e87451c025ab101d2e56f751e7"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v1.1.0/proto_1.1.0_linux_armv6.zip"
      sha256 "2d0bfedf40282e1e0caa344c9e09c8541dc67a80ffb8764d58383d43044557dd"

      def install
        bin.install "proto"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/BitsOfAByte/proto/releases/download/v1.1.0/proto_1.1.0_linux_arm64.zip"
      sha256 "ccc740f6d5a46dc537f8fff03fa5b428d36e768f7d71a8d545eb52b44958690d"

      def install
        bin.install "proto"
      end
    end
  end
end
