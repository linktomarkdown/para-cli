# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class V010 < Formula
  desc "ParaSnack模板生成脚手架"
  homepage "https://github.com/linktomarkdown/para-cli"
  url "https://github.com/linktomarkdown/para-cli/releases/download/v0.1.0/para_darwin_amd64"
  sha256 "7c20580f91aa87068cce7da7664f55052cf475f28c131bd72d683bb34402b521"
  license ""

  # depends_on "cmake" => :build
  # depends_on "go" => :build
  def install
    # ENV.deparallelize  # if your formula fails when building in parallel
    # Remove unrecognized options if warned by configure
    # https://rubydoc.brew.sh/Formula.html#std_configure_args-instance_method
    system "./configure", *std_configure_args, "--disable-silent-rules"
    # system "cmake", "-S", ".", "-B", "build", *std_cmake_args
    # system "go", "build", "-o", bin/"para"
    bin.install "para_darwin_amd64" => "para"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test v0.1.0`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
