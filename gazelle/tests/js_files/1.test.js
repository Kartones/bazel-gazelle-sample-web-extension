import assert from "node:assert";
import { describe, it } from "node:test";
import { value } from "./1.js";

describe("Tests for 1", () => {
  it("Checks 1's value", () => {
    assert.strictEqual(value, "1");
  });
});
