const wasm = require('wasm-bindgen-test.js');
const assert = require('assert');

exports.js_c_style_enum = () => {
    assert.strictEqual(wasm.Color.Green, 0);
    assert.strictEqual(wasm.Color.Yellow, 1);
    assert.strictEqual(wasm.Color.Red, 2);
    assert.strictEqual(Object.keys(wasm.Color).length, 3);

    assert.strictEqual(wasm.enum_cycle(wasm.Color.Green), wasm.Color.Yellow);
};

exports.js_c_style_enum_with_custom_values = () => {
    assert.strictEqual(wasm.ColorWithCustomValues.Green, 21);
    assert.strictEqual(wasm.ColorWithCustomValues.Yellow, 34);
    assert.strictEqual(wasm.ColorWithCustomValues.Red, 2);
    assert.strictEqual(Object.keys(wasm.ColorWithCustomValues).length, 3);

    assert.strictEqual(wasm.enum_with_custom_values_cycle(wasm.ColorWithCustomValues.Green), wasm.ColorWithCustomValues.Yellow);
};
