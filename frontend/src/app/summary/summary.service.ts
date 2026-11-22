import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { WorkflowSummary } from './workflow-summary';

@Injectable({
  providedIn: 'root',
})
export class SummaryService {
  private readonly http = inject(HttpClient);

  getSummary(): Observable<WorkflowSummary> {
    return this.http.get<WorkflowSummary>('/api/v1/summary');
  }
}
