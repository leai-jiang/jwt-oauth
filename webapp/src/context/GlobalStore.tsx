import * as React from "react";
import store from "./store";

const GlobalStore = React.createContext(store);

export default GlobalStore;
