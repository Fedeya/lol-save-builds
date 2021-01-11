import create from 'zustand';
import { Champion } from '../interfaces';

type State = {
  builds: Champion[];
  loadBuilds(): Promise<void>;
  addBuild(url: string): Promise<void>;
};

export const useBuildsStore = create<State>((set, get) => ({
  builds: [],
  async loadBuilds() {
    const builds = await window.backend.getBuilds();

    set({ builds });
  },
  async addBuild(url) {
    try {
      const champion = await window.backend.scrap(url);
      const builds = get().builds;
      set({ builds: [...builds, champion] });
    } catch (err) {
      console.log(err);
    }
  }
}));
