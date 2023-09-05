import { SnackbarPosition } from '@shared/components/snackbar/models/snackbar-position.model';
import { SnackbarType } from '@shared/components/snackbar/models/snackbar-type.model';

export interface SnackbarConfig {
  message: string;
  title?: string;
  type?: SnackbarType;
  duration?: number;
  buttonLabel?: string;
  position?: SnackbarPosition;
}
