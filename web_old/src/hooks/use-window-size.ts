import { useEffect, useState } from "react";

export const useWindowSize = (): {width: number, height: number} => {
  const [windowSize, setWindowSize] = useState<{width: number, height: number}>({
      width: window.innerWidth,
      height: window.innerHeight
  });

  useEffect(() => {
    const handleResize = () => {
      setWindowSize({
        width: window.innerWidth,
        height: window.innerHeight,
      })
    }

    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  return windowSize;
}
