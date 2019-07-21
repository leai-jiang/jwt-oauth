import Home from "./routes/Home";

interface Route {
    key: string;
    path: string;
    exact?: boolean;
    auth?: 0 | 1;
    component: ((props?: any) => JSX.Element) | any;
}

export const routes: Route[] = [
    {
      key: "home",
      path: "/home",
      component: Home
    }
];
