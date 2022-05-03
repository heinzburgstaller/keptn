import { Target } from '../../../shared/interfaces/indicator-result';
import { ResultTypes } from '../../../shared/models/result-types';

export enum IHeatmapTooltipType {
  SCORE,
  SLI,
}

export enum EvaluationResultTypeExtension {
  INFO = 'info',
}
export type EvaluationResultType = ResultTypes | EvaluationResultTypeExtension;

export interface IHeatmapScoreTooltip {
  type: IHeatmapTooltipType.SCORE;
  passCount: number;
  warningCount: number;
  failedCount: number;
  thresholdPass: number;
  thresholdWarn: number;
  fail: boolean;
  warn: boolean;
  value: number;
}

export interface IHeatmapSliTooltip {
  type: IHeatmapTooltipType.SLI;
  keySli: boolean;
  score: number;
  passTargets: Target[];
  warningTargets: Target[];
  value: number;
}

export type IHeatmapTooltip = IHeatmapScoreTooltip | IHeatmapSliTooltip;

export interface IDataPoint {
  date: string;
  sli: string;
  tooltip: IHeatmapTooltip;
  color: EvaluationResultType;
  comparedIndices: number[];
  /**
   * Unique identifier like keptnContext that can be used on tileSelected
   */
  identifier: string;
}
