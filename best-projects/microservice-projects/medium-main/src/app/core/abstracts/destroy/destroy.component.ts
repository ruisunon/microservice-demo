import { Component, OnDestroy } from '@angular/core';
import { Subject } from 'rxjs';

@Component({ selector: 'app-destroy', standalone: true, template: `` })
export abstract class DestroyComponent implements OnDestroy {
  protected readonly destroy$: Subject<void> = new Subject<void>();

  public ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
