interface Backend {
  scrap(url: string): Promise<void>;
  getBuilds(): Promise<any>;
}

export declare global {
  interface Window {
    backend: Backend;
  }
}
