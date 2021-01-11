import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import { Champion } from '../interfaces';
import { useBuildsStore } from '../store/builds';

const Build: React.FC = () => {
  const [build, setBuild] = useState<Champion>();
  const builds = useBuildsStore(state => state.builds);
  const { id } = useParams<{ id: string }>();

  useEffect(() => {
    const build = builds.find(build => build.id === id);
    console.log(build);
    setBuild(build);
  }, [builds, id]);

  if (!build) return <p>Build Not Found</p>;

  return (
    <div className="p-5">
      <div className="flex flex-col">
        <div className="flex flex-row">
          <img src={build.image_url} className="mr-3" alt={build.name} />
          <div className="flex flex-col">
            <h1>{build.name}</h1>
            <h2>{build.line}</h2>
          </div>
        </div>

        <div className="flex flex-col justify-center">
          <h2>Runes</h2>
          <div className="flex">
            {build.runes.map((rune, i) => (
              <div key={i} className="flex mr-3">
                <div className="flex flex-col">
                  <img
                    src={rune.primary.type.image_url}
                    alt={rune.primary.type.name}
                    className="w-10 h-10 mr-3"
                  />
                  <img
                    src={rune.primary.main.image_url}
                    alt={rune.primary.main.name}
                    className="w-10 h-10 mr-3"
                  />
                  {rune.primary.more.map((more, i) => (
                    <img
                      key={i}
                      className="w-10 h-10 mr-3"
                      src={more.image_url}
                      alt={more.name}
                    />
                  ))}
                </div>
                <div className="flex flex-col">
                  <img
                    src={rune.secondary.type.image_url}
                    alt={rune.secondary.type.name}
                    className="w-10 h-10 mr-3"
                  />
                  {rune.secondary.more.map((more, i) => (
                    <img
                      key={i}
                      className="w-10 h-10 mr-3"
                      src={more.image_url}
                      alt={more.name}
                    />
                  ))}
                </div>
                <div className="flex flex-col">
                  {rune.boosts.map((boost, i) => (
                    <img
                      key={i}
                      src={boost.image_url}
                      alt={boost.name}
                      className="w-6 h-6"
                    />
                  ))}
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Build;
