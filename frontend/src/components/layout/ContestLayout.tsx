import type { ReactNode } from "react";
import { ContestSidebar } from "../model/contest/ContestSidebar";
import { GLOBAL_HEADER_H } from "../ui/GlobalHeader";
import { WithHeaderFooter } from "./WithHeaderFooter";

export type ContestLayoutProps = {
  children: ReactNode;
};

export const ContestLayout = ({ children }: ContestLayoutProps) => {
  return (
    <WithHeaderFooter
      headerProps={{ contestSlug: "sbc001", contestTitle: "SZPP Beginners Contest 001" }}
      leftChildren={
        <ContestSidebar
          top={GLOBAL_HEADER_H}
          startAt={new Date("2023-09-02 21:00")}
          endAt={new Date("2023-09-04 01:00")}
          now={new Date("2023-09-02 21:00:01")}
          slug="sbc001"
          tasks={[
            { id: 1000, title: "すずっぴー君のおつかい" },
            { id: 1001, title: "すずっぴー君の変死", scoreStatus: "perfect" },
            { id: 1002, title: "すずっぴー君の怪死", scoreStatus: "perfect" },
            { id: 1003, title: "すずっぴー君の爆死", scoreStatus: "zero" },
            { id: 1004, title: "すずっぴー君の溺死" },
            { id: 1005, title: "すずっぴー君の圧死", scoreStatus: "partial" },
            { id: 1006, title: "すずっぴー君の尊死", scoreStatus: "zero" },
            { id: 1007, title: "長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル長いタイトル" },
            { id: 1008, title: "longlonglonglonglonglonglonglonglonglonglonglong", scoreStatus: "partial" },
            { id: 1009, title: "すずっぴー君のおつかい" },
          ]}
        />
      }
    >
      {children}
    </WithHeaderFooter>
  );
};
