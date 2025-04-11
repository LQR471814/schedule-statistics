import { instantToTimestamp } from "$lib/time";
import { Temporal } from "@js-temporal/polyfill";
import { client } from "./rpc";
import { toast } from "svelte-sonner"
import type { EventsResponse } from "$api/api_pb";

export type Interval = {
	start: Temporal.ZonedDateTime;
	end: Temporal.ZonedDateTime;
};

export enum IntervalOption {
	THIS_DAY = 0,
	THIS_WEEK = 1,
	THIS_MONTH = 2,
	THIS_YEAR = 3,
	LAST_3_MONTHS = 4,
	LAST_6_MONTHS = 5,
	CUSTOM = 6,
}

export class EventState {
	option = $state(IntervalOption.THIS_WEEK)
	customBounds: Interval = $state<Interval>() as Interval
	events = $state.raw<EventsResponse>()

	private interval: Interval = $derived.by((): Interval => {
		const now = Temporal.Now.zonedDateTimeISO();

		switch (this.option) {
			case IntervalOption.CUSTOM:
				return this.customBounds;
			case IntervalOption.THIS_DAY: {
				const start = now.subtract({
					hours: now.hour,
					minutes: now.minute,
					seconds: now.second,
					milliseconds: now.millisecond,
					nanoseconds: now.nanosecond,
				});
				const end = start.add({
					hours: 23,
					minutes: 59,
					seconds: 59,
					milliseconds: 999,
					nanoseconds: 999,
				});
				return { start, end };
			}
			case IntervalOption.THIS_WEEK: {
				const start = now.subtract({
					days: now.dayOfWeek - 1,
					hours: now.hour,
					minutes: now.minute,
					seconds: now.second,
					milliseconds: now.millisecond,
					nanoseconds: now.nanosecond,
				});
				const end = start.add({
					days: now.daysInWeek,
				});
				return { start, end };
			}
			case IntervalOption.THIS_MONTH: {
				const start = now.subtract({
					days: now.day - 1,
					hours: now.hour,
					minutes: now.minute,
					seconds: now.second,
					milliseconds: now.millisecond,
					nanoseconds: now.nanosecond,
				});
				const end = start.add({
					months: 1,
				});
				return { start, end };
			}
			case IntervalOption.THIS_YEAR: {
				const start = now.subtract({
					days: now.dayOfYear,
					hours: now.hour,
					minutes: now.minute,
					seconds: now.second,
					milliseconds: now.millisecond,
					nanoseconds: now.nanosecond,
				});
				const end = start.add({
					years: 1,
				});
				return { start, end };
			}
			case IntervalOption.LAST_3_MONTHS: {
				const start = now.subtract({
					months: 3,
					hours: now.hour,
					minutes: now.minute,
					seconds: now.second,
					milliseconds: now.millisecond,
					nanoseconds: now.nanosecond,
				});
				const end = now;
				return { start, end };
			}
			case IntervalOption.LAST_6_MONTHS: {
				const start = now.subtract({
					months: 6,
					hours: now.hour,
					minutes: now.minute,
					seconds: now.second,
					milliseconds: now.millisecond,
					nanoseconds: now.nanosecond,
				});
				const end = now;
				return { start, end };
			}
		}
	})

	constructor() {
		const now = Temporal.Now.zonedDateTimeISO();
		this.customBounds = {
			start: now.subtract({ weeks: 1 }),
			end: now.add({ weeks: 1 }),
		}

		$effect(() => {
			toast.promise(
				() => client.events({
					interval: {
						start: instantToTimestamp(this.interval.start.toInstant()),
						end: instantToTimestamp(this.interval.end.toInstant()),
					},
				})
					.then((res) => { this.events = res }),
				{
					loading: 'Fetching events...',
					success: 'Fetch events: Success',
					error: 'Fetch events: Error',
					dismissable: true,
					duration: 500,
				}
			)
		})
	}
}

