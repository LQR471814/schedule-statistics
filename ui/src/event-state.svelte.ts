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
	THIS_DAY = "THIS_DAY",
	THIS_WEEK = "THIS_WEEK",
	THIS_MONTH = "THIS_MONTH",
	THIS_YEAR = "THIS_YEAR",
	LAST_3_MONTHS = "LAST_3_MONTHS",
	LAST_6_MONTHS = "LAST_6_MONTHS",
	CUSTOM = "CUSTOM",
}

const full_day = {
	hours: 23,
	minutes: 59,
	seconds: 59,
	milliseconds: 999,
	nanoseconds: 999,
}

export class EventModel {
	version = $state(0)
	option = $state(IntervalOption.THIS_WEEK)
	customBounds: Interval = $state<Interval>() as Interval
	events = $state.raw<EventsResponse>()

	interval: Interval = $derived.by((): Interval => {
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
				const end = start.add(full_day);
				return { start, end };
			}
			case IntervalOption.THIS_WEEK: {
				const start = now.subtract({
					days: now.dayOfWeek,
					hours: now.hour,
					minutes: now.minute,
					seconds: now.second,
					milliseconds: now.millisecond,
					nanoseconds: now.nanosecond,
				});
				const end = start.add({
					days: now.daysInWeek - 1,
					...full_day,
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
					...full_day,
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
					...full_day,
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
			this.interval
			this.refresh()
		})
	}

	refresh(): Promise<void> {
		return new Promise((resolve, reject) => {
			toast.promise(
				() => client.events({
					timezone: Temporal.Now.timeZoneId(),
					interval: {
						start: instantToTimestamp(this.interval.start.toInstant()),
						end: instantToTimestamp(this.interval.end.toInstant()),
					},
				})
					.then((res) => {
						this.events = res
						console.table(
							res.events.map((e) => {
								const startTime = Temporal.Instant
									.fromEpochMilliseconds(Number(e.interval!.start!.seconds) * 1000)
									.toZonedDateTimeISO(Temporal.Now.timeZoneId())

								const endTime = Temporal.Instant
									.fromEpochMilliseconds(Number(e.interval!.end!.seconds) * 1000)
									.toZonedDateTimeISO(Temporal.Now.timeZoneId())

								return {
									name: res.eventNames[e.name],
									tag: e.tags.map((t) => res.tags[t])[0],
									startTime: `${startTime.year}-${startTime.month}-${startTime.day} ${startTime.hour}`,
									endTime: `${endTime.year}-${endTime.month}-${endTime.day} ${endTime.hour}`
								}
							})
						)
						resolve()
					})
					.catch((err) => {
						reject(err)
						throw err
					}),
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

