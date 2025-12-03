import fs from 'fs';
import path from 'path';
import sharp from 'sharp';

const root = process.cwd();
const baseDir = path.join(root, 'public', 'img');
const exts = new Set(['.png', '.jpg', '.jpeg', '.webp']);

async function optimizeFile(filePath) {
  const ext = path.extname(filePath).toLowerCase();
  try {
    const input = fs.readFileSync(filePath);
    let output;
    if (ext === '.webp') {
      output = await sharp(input).webp({ quality: 75 }).toBuffer();
    } else if (ext === '.png') {
      output = await sharp(input).png({ compressionLevel: 9 }).toBuffer();
    } else if (ext === '.jpg' || ext === '.jpeg') {
      output = await sharp(input).jpeg({ quality: 80, mozjpeg: true }).toBuffer();
    } else {
      return { skipped: true };
    }
    if (output.length < input.length) {
      fs.writeFileSync(filePath, output);
      return { saved: input.length - output.length, before: input.length, after: output.length };
    }
    return { saved: 0, before: input.length, after: output.length };
  } catch (e) {
    return { error: e.message };
  }
}

function* walk(dir) {
  const entries = fs.readdirSync(dir, { withFileTypes: true });
  for (const entry of entries) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      yield* walk(full);
    } else {
      yield full;
    }
  }
}

async function main() {
  if (!fs.existsSync(baseDir)) {
    console.log('No public/img directory found');
    process.exit(0);
  }
  let totalSaved = 0;
  let processed = 0;
  let skipped = 0;
  let failed = 0;
  for (const file of walk(baseDir)) {
    const ext = path.extname(file).toLowerCase();
    if (!exts.has(ext)) continue;
    const res = await optimizeFile(file);
    if (res.error) { failed++; continue; }
    if (res.skipped) { skipped++; continue; }
    processed++;
    totalSaved += res.saved || 0;
  }
  console.log(`Optimización completada. Procesados: ${processed}, Omitidos: ${skipped}, Fallidos: ${failed}, Ahorro total: ${(totalSaved/1048576).toFixed(2)} MB`);
}

main();


