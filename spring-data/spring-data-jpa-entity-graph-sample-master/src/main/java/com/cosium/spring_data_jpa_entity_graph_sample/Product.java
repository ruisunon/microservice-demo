package com.cosium.spring_data_jpa_entity_graph_sample;

import javax.persistence.Access;
import javax.persistence.AccessType;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.NamedAttributeNode;
import javax.persistence.NamedEntityGraph;
import javax.persistence.NamedEntityGraphs;

/**
 * @author Réda Housni Alaoui
 */
@NamedEntityGraphs(@NamedEntityGraph(name = Product.WITH_BRAND_EG, attributeNodes = @NamedAttributeNode(Product_.BRAND)))
@Entity
public class Product {

	public static final String WITH_BRAND_EG = "product(brand)";

	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	@Access(value = AccessType.PROPERTY)
	private long id = 0;

	private String name;

	@ManyToOne(fetch = FetchType.LAZY, optional = false)
	@JoinColumn(nullable = false)
	private Brand brand;

	public Product() {

	}

	public Product(Brand brand, String name) {
		this.name = name;
		this.brand = brand;
	}

	public void setId(long id) {
		this.id = id;
	}

	public long getId() {
		return id;
	}

	public String getName() {
		return name;
	}

	public Brand getBrand() {
		return brand;
	}
}
