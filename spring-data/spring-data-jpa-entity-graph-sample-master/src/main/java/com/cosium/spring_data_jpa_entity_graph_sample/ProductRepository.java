package com.cosium.spring_data_jpa_entity_graph_sample;

import com.cosium.spring.data.jpa.entity.graph.repository.EntityGraphJpaRepository;
import com.cosium.spring.data.jpa.entity.graph.repository.EntityGraphJpaSpecificationExecutor;

/**
 * @author Réda Housni Alaoui
 */
public interface ProductRepository extends EntityGraphJpaRepository<Product, Long>, EntityGraphJpaSpecificationExecutor<Product> {
}
