package com.cosium.spring_data_jpa_entity_graph_sample;

import com.cosium.spring.data.jpa.entity.graph.domain2.NamedEntityGraph;
import java.util.stream.Collectors;
import java.util.stream.StreamSupport;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

/**
 * @author Réda Housni Alaoui
 */
@Controller
@RequestMapping("/products")
public class ProductController {

	private final ProductRepository repository;

	public ProductController(ProductRepository repository) {
		this.repository = repository;
	}

	@GetMapping
	public ResponseEntity<?> list(@RequestParam(name = "with_brand", required = false, defaultValue = "false") boolean withBrand) {
		Iterable<Product> products;
		if (withBrand) {
			products = repository.findAll(NamedEntityGraph.loading(Product.WITH_BRAND_EG));
		} else {
			products = repository.findAll();
		}
		return ResponseEntity.ok(StreamSupport.stream(products.spliterator(), false)
				.map(product -> new ProductRepresentation(product, withBrand))
				.collect(Collectors.toList()));
	}

	@GetMapping("/{id}")
	public ResponseEntity<?> findById(@PathVariable("id") long id) {
		return ResponseEntity.ok(
				repository.findById(
						id,
						ProductEntityGraph
								.____()
								.brand()
								.____
								.____()
				).map(product -> new ProductRepresentation(product, true))
		);
	}

	private static class ProductRepresentation {
		private final long id;
		private final String name;
		private final String brandName;

		ProductRepresentation(Product product, boolean withBrand) {
			id = product.getId();
			name = product.getName();
			if (withBrand) {
				brandName = product.getBrand().getName();
			} else {
				brandName = "Not loaded";
			}
		}

		public long getId() {
			return id;
		}

		public String getName() {
			return name;
		}

		public String getBrandName() {
			return brandName;
		}
	}
}
