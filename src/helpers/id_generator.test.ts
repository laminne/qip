import { assert, describe, it } from "vitest";
import { SnowflakeIDGenerator } from "./id_generator.js";

const generator = new SnowflakeIDGenerator(0);

describe("ID生成ができる", () => {
  it("複数回生成してもIDが重複しない", () => {
    let oldID = "";
    for (let i = 0; i < 10000; i++) {
      const newID = generator.generate();
      assert.notEqual(newID, oldID);
      oldID = newID;
    }
  });
});
