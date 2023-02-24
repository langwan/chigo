module.exports = function override(config, env) {
  console.log("override");
  let loaders = config.resolve;
  loaders.fallback = {
    fs: false,
    tls: false,
    net: false,
    stream: require.resolve("stream-browserify"),
    crypto: require.resolve("crypto-browserify"),
  };

  return config;
};
