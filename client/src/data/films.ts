import { randInt } from "../random";

const filmsDataPath = "films.gz";

export type Film = {
  title: string;
  year: string;
  actor: string;
  plot: string[];
};
export type Slug = string;
export type FilmsBySlug = { [slug: Slug]: Film };
type FilmWithSlug = { slug: Slug; film: Film };

var _films: FilmsBySlug = {};
var _filmsYearAsc: FilmWithSlug[] = [];

export async function films(): Promise<FilmsBySlug> {
  if (Object.keys(_films).length === 0) {
    const resp = await fetch(filmsDataPath);
    const json = await resp.text();
    _films = JSON.parse(json);
  }
  const f = Object.entries(_films);
  return _films;
}

async function byYearAsc(): Promise<FilmWithSlug[]> {
  if (_filmsYearAsc.length === 0) {
    const f = await films();
    _filmsYearAsc = Object.entries(f).map(([slug, film]) => ({ slug, film }));
    _filmsYearAsc.sort((a, b) => parseInt(a.film.year) - parseInt(b.film.year));
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
    result[f.slug] = f.film;
  }
  return result;
}
