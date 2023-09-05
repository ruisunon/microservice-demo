import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { ComponentRef, Injectable, Injector, inject } from '@angular/core';
import { SNACKBAR_CONFIG } from '@shared/components/snackbar/constants/snackbar-config-token';
import { SnackbarConfig } from '@shared/components/snackbar/models/snackbar-config.model';
import { SnackbarPosition } from '@shared/components/snackbar/models/snackbar-position.model';
import { SnackbarComponent } from '@shared/components/snackbar/snackbar.component';

@Injectable({ providedIn: 'root' })
export class SnackbarService {
  private readonly overlayRef: OverlayRef = inject(Overlay).create();
  private readonly injector: Injector = inject(Injector);

  private componentRef!: ComponentRef<SnackbarComponent>;

  public openSnackbar(config: SnackbarConfig): void {
    this.componentRef = this.overlayRef.attach<SnackbarComponent>(
      new ComponentPortal(SnackbarComponent, null, this.createInjector(config))
    );

    config.duration &&
      setTimeout((): void => {
        this.destroySnackbarComponent();
      }, config.duration);
  }

  public close(): void {
    this.componentRef && this.destroySnackbarComponent();
  }

  public setSnackbarPosition(position?: SnackbarPosition): string {
    switch (position) {
      case 'top-right':
        return 'snackbar--top-right';
      case 'top-left':
        return 'snackbar--top-left';
      case 'bottom-center':
        return 'snackbar--bottom-center';
      case 'bottom-left':
        return 'snackbar--bottom-left';
      case 'bottom-right':
        return 'snackbar--bottom-right';
      default:
        return 'snackbar--top-center';
    }
  }

  private createInjector(config: SnackbarConfig): Injector {
    return Injector.create({
      parent: this.injector,
      providers: [{ provide: SNACKBAR_CONFIG, useValue: config }],
    });
  }

  private destroySnackbarComponent(): void {
    this.overlayRef.detach();
    this.componentRef.destroy();
  }
}
