package com.cosium.spring_data_jpa_entity_graph_sample;

import org.springframework.context.SmartLifecycle;
import org.springframework.stereotype.Component;

/**
 * @author Réda Housni Alaoui
 */
@Component
public class DatabaseInitializer implements SmartLifecycle {

	private final BrandRepository brandRepository;
	private final ProductRepository productRepository;
	private boolean running;

	public DatabaseInitializer(BrandRepository brandRepository, ProductRepository productRepository) {
		this.brandRepository = brandRepository;
		this.productRepository = productRepository;
	}

	@Override
	public void start() {
		Brand tesla = brandRepository.save(new Brand("Tesla"));
		productRepository.save(new Product(tesla, "Model S"));
		productRepository.save(new Product(tesla, "Model Y"));
		productRepository.save(new Product(tesla, "Model X"));

		Brand audi = brandRepository.save(new Brand("Audi"));
		productRepository.save(new Product(audi, "A7"));
		productRepository.save(new Product(audi, "TT"));
		productRepository.save(new Product(audi, "R8"));

		running = true;
	}

	@Override
	public void stop() {
		running = false;
	}

	@Override
	public boolean isRunning() {
		return running;
	}
}
