package com.cosium.spring_data_jpa_entity_graph_sample;

import javax.persistence.Access;
import javax.persistence.AccessType;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

/**
 * @author Réda Housni Alaoui
 */
@Entity
public class Brand {

	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	@Access(value = AccessType.PROPERTY)
	private long id = 0;

	private String name;

	public Brand(){

	}

	public Brand(String name) {
		this.name = name;
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
}
