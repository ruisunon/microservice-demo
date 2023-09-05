import { CommonModule } from '@angular/common';
import { ChangeDetectionStrategy, Component, Input, OnInit } from '@angular/core';
import { RouterLink } from '@angular/router';
import { getRange } from '@core/utils/get-range';

const PaginatorImports: Array<any> = [CommonModule, RouterLink];

@Component({
  selector: 'app-paginator',
  standalone: true,
  imports: PaginatorImports,
  templateUrl: './paginator.component.html',
  styleUrls: ['./paginator.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PaginatorComponent implements OnInit {
  @Input() public total: number = 0;
  @Input() public limit: number = 20;
  @Input() public currentPage: number = 1;
  @Input() public url: string = '';

  public pagesCount: number = 0;
  public pages: number[] = [];

  public ngOnInit(): void {
    this.pagesCount = Math.ceil(this.total / this.limit);
    this.pages = this.pagesCount > 0 ? getRange(1, this.pagesCount) : [];
  }
}
