module.exports = function (api) {
    api.cache(true);
    const presets = [
        ["@babel/preset-env", { "modules": false }],
        "@babel/preset-react"
    ];
    const plugins = [
        "@babel/plugin-transform-runtime",
        "@babel/plugin-transform-flow-strip-types",
        "@babel/plugin-proposal-class-properties"
    ];

    return {
        presets,
        plugins
    };
}