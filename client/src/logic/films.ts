import { filmsDataPath } from "../config";
import { randInt } from "../random";
import { Film, Slug, FilmsBySlug } from "../types";

var _films: FilmsBySlug = {};
var _filmsYearAsc: Film[] = [];

export async function films(): Promise<FilmsBySlug> {
  if (Object.keys(_films).length === 0) {
    const resp = await fetch(filmsDataPath);
    const json = await resp.text();
    // structure is {slug: {year: 1984, title: ...}, ...}
    const parsed: { [s: Slug]: Omit<Film, "slug"> } = JSON.parse(json);
    for (const [slug, film] of Object.entries(parsed)) {
      _films[slug] = { slug, ...film };
    }
  }
  return _films;
}

async function byYearAsc(): Promise<Film[]> {
  if (_filmsYearAsc.length === 0) {
    _filmsYearAsc = Object.values(await films());
    _filmsYearAsc.sort((a, b) => parseInt(a.year) - parseInt(b.year));
  }
  return _filmsYearAsc;
}

// Pick N films from the list, contiguous by year.
export async function pick(n: number): Promise<FilmsBySlug> {
  const all = await byYearAsc();
  const i = randInt(0, all.length - n);
  const fs = all.slice(i, i + n);
  const result: FilmsBySlug = {};
  for (const f of fs) {
    result[f.slug] = f;
  }
  return result;
}
