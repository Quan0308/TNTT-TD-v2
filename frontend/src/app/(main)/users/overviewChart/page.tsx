import AreaChartOverview from "@/app/(main)/users/overviewChart/Charts/AreaChart";
import BarChartOverview from "@/app/(main)/users/overviewChart/Charts/BarChart";

export default function OverviewChart() {
  return (
    <section className="p-3 pr-10 xl:pr-3">
      <div className="w-full h-fit flex flex-col gap-2 items-center">
        <AreaChartOverview />
        <BarChartOverview />
      </div>
    </section>
  );
}
