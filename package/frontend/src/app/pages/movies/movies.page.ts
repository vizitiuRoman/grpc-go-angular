import { Component, OnInit } from '@angular/core';

import { MovieGrpcService } from '@grpc/services/movie/movie.service';

@Component({
    selector: 'app-movies',
    templateUrl: './movies.page.html',
    styleUrls: ['./movies.page.scss'],
})
export class MoviesPage implements OnInit {
    constructor(
        private movieService: MovieGrpcService
    ) {
    }

    ngOnInit(): void {
        this.movieService.getMovies().subscribe();
    }

    public doInfinite($event: Event): void {
    }
}
