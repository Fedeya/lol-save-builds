import create from 'zustand';

type State = {
  builds: any[];
};

export const useBuildsStore = create<State>(() => ({
  builds: []
}));
