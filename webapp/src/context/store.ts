interface State {
  isSign: boolean;
  userName?: string;
  avatar?: string;
  token?: string;
}

interface Store {
  state: State;
  dispatch: Function;
}

const store: Store = {
  state: {
    isSign: false
  },
  dispatch: () => {}
};

export { State, Store, store as default };
