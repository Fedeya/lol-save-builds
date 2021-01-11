import { Champion } from './interfaces';

interface Backend {
  scrap(url: string): Promise<Champion>;
  getBuilds(): Promise<Champion[]>;
}

export declare global {
  interface Window {
    backend: Backend;
  }
}
