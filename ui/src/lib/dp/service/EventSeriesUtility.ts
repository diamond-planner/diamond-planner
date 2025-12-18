import type {EventSeriesState} from "$lib/dp/types/EventSeriesState.ts";

export class EventSeriesUtility {
  public static getSeriesState(startDate: Date, endDate: Date): EventSeriesState {
    const nowDate = new Date();

    if (nowDate < startDate) {
      return "future";
    }
    if (nowDate > endDate) {
      return "past";
    }
    return "ongoing";
  }
}
