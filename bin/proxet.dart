import 'dart:io';

void main(List<String> args) {
  if (args.isNotEmpty) {
    switch (args[0]) {
      case '-r':
        resetProxy();
        return;
      case '-a':
        if (args.length < 2) {
          print("\x1B[31mPlease provide a proxy address\x1B[0m");
          return;
        }
        setProxy(args[1]);
        return;
      case '-c':
        if (isProxySet()) {
          print("\x1B[32mProxy is set\x1B[0m");
        } else {
          print("\x1B[31mProxy is not set\x1B[0m");
        }
        return;
    }
  }
  print("Usage:");
  print("proxet -a <proxy_address> to set the proxy");
  print("proxet -r to reset the proxy settings");
  print("proxet -c to check if the proxy is set");
}

final String configPath =
    '${Platform.environment['HOME']}/.config/fish/config.fish';

void setProxy(String address) {
  if (isProxySet()) {
    print("\x1B[33mProxy is already set\x1B[0m");
    return;
  }

  final file = File(configPath);
  if (!file.existsSync()) {
    print("\x1B[31mCouldn't open fish config file\x1B[0m");
    return;
  }

  final proxySettings = '''

set -gx http_proxy "$address"
set -gx https_proxy "$address"
set -gx ftp_proxy "$address"
set -gx all_proxy "$address"
set -gx no_proxy "$address"
''';

  file.writeAsStringSync(proxySettings, mode: FileMode.append);
  sourceConfig();
  print("\x1B[32mProxy settings updated!\x1B[0m");
}

void resetProxy() {
  final file = File(configPath);
  if (!file.existsSync()) {
    print("\x1B[31mCouldn't open fish config file\x1B[0m");
    return;
  }

  final lines =
      file.readAsLinesSync().where((line) => !line.contains("_proxy")).toList();
  file.writeAsStringSync(lines.join('\n'));

  sourceConfig();
  print("\x1B[32mProxy settings reset!\x1B[0m");
}

bool isProxySet() {
  final file = File(configPath);
  if (!file.existsSync()) {
    print("\x1B[31mCouldn't open fish config file\x1B[0m");
    return false;
  }

  return file.readAsLinesSync().any(
    (line) =>
        line.contains("http_proxy") ||
        line.contains("https_proxy") ||
        line.contains("ftp_proxy") ||
        line.contains("all_proxy") ||
        line.contains("no_proxy"),
  );
}

void sourceConfig() {
  Process.runSync("fish", ["-c", "source $configPath"]);
}
