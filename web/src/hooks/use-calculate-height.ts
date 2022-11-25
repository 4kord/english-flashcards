import { useEffect, useState } from "react";
import { useWindowSize } from "@/hooks/use-window-size";

export const useCalculateHeight = (): number | null => {
  const windowSize = useWindowSize();

  const [height, setHeight] = useState<null | number>(null);

  useEffect(() => {
    setHeight(windowSize.height - document.getElementById("33m5WuJe")?.clientHeight!);
  }, [windowSize])

  return height
}
