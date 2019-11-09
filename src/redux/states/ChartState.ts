// import { DisplayRate } from "../interfaces/Rates";
import { Action, AnyAction, Dispatch } from "redux";
// import { ThunkAction } from "redux-thunk";
// import { actionCreator, AppState } from "../../store";
// import { firestoreApp } from "../../firebase";
//
//
// export type RatePayload = {
//   rates: Array<DisplayRate>
// };
//
//
// export interface RateState {
//   rates: Array<DisplayRate>,
// }
//
// export const initialState: RateState = {
//   rates: []
// };
//
// //Actionの定義
// //Action-Creatorの定義
// //Reducerの定義
// export interface LoadedRateAction extends Action {
//   type: "LOADED_RATE";
//   payload: RatePayload;
// }
//
// export const loadedRate = (payload: RatePayload): LoadedRateAction => {
//   return {
//     type: "LOADED_RATE",
//     payload
//   };
// };
//
// export const rateStateReducer = (state: RateState = initialState, action: LoadedRateAction) => {
//   switch (action.type) {
//     case "LOADED_RATE": {
//       const rates = action.payload.rates;
//       return { ...state, rates };
//     }
//     default:
//       return state;
//   }
// };
//
// //Thunk-Actionの定義
// export const firestoreRateConnect = (payload: void): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
//   firestoreApp.collection("rates")
//     .orderBy("sendAt")
//     .onSnapshot((snapshots) => {
//       const docs = snapshots.docs.map((snapshot => {
//         const data = snapshot.data();
//         return { name: data["sendAt"], uv: data["ppm"], amt: 20 };
//       }));
//       const payload: RatePayload = ({ rates: docs });
//       dispatch(actionCreator.rates.loadedRate(payload));
//     });
// };
//
// //ドメイン毎にまとめる。
// export const rateActionCreator = {
//   loadedRate,
//   firestoreRateConnect
// };