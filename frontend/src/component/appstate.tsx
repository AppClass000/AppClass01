import { createContext } from "react";

interface menuContextType {
  isOpened: boolean;
  setOpened: (value: boolean) => void;
}

const menuContext = createContext<menuContextType>({
  isOpened: false,
  setOpened: () => {},
});


export default menuContext;

