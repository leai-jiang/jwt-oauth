import { State } from "./store";

interface Action {
  type: string;
  payload: any;
}

function reducer(state: State, action: Action) {
  switch (action.type) {
    case "SIGN_IN":
      return {
        ...state,
        isSign: true,
        ...action.payload
      };
    case "SIGN_OUT":
      return {
        isSign: false,
      };
    default:
      return state;
  }
}

export default reducer;
