const fs = require("fs");

for (let i = 0; i < 1e2; i++) {
  const randomLen = Math.floor(Math.random() * 100);
  for (let j = 0; j < randomLen; j++) {
    fs.appendFileSync(
      `./measurements.txt`,
      `measurement${i};${Math.random() * 1000}\n`
    );
  }
}
