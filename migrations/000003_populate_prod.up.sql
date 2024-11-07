TRUNCATE products CASCADE;

-- Função para gerar produtos em massa
DO $$
DECLARE
    -- Arrays com palavras para gerar combinações
    nomes_comida TEXT[] := ARRAY['Pizza', 'Hambúrguer', 'Sushi', 'Salada', 'Pasta', 'Sanduíche', 'Sobremesa', 'Prato', 'Porção', 'Combo'];
    sabores TEXT[] := ARRAY['Frango', 'Carne', 'Vegano', 'Vegetariano', 'Especial', 'Premium', 'Tradicional', 'Gourmet', 'Light', 'Fit'];
    complementos TEXT[] := ARRAY['com Queijo', 'ao Molho', 'Grelhado', 'Crocante', 'Supremo', 'da Casa', 'Premium', 'Extra', 'Super', 'Master'];
    lojas TEXT[] := ARRAY['store1', 'store2', 'store3', 'store4', 'store5'];

    -- Variáveis para o loop
    nome TEXT;
    descricao TEXT;
    preco DECIMAL;
    quantidade INT;
    loja TEXT;
    i INT;
BEGIN
    -- Loop para gerar 100.000 produtos
    FOR i IN 1..100000 LOOP
        -- Seleciona aleatoriamente elementos dos arrays
        nome := nomes_comida[1 + floor(random() * array_length(nomes_comida, 1))] || ' ' ||
                sabores[1 + floor(random() * array_length(sabores, 1))] || ' ' ||
                complementos[1 + floor(random() * array_length(complementos, 1))];

        -- Gera uma descrição mais longa
        descricao := 'Delicioso ' || lower(nome) || ' preparado com ingredientes selecionados. ' ||
                    'Acompanha ' || sabores[1 + floor(random() * array_length(sabores, 1))] ||
                    ' e ' || complementos[1 + floor(random() * array_length(complementos, 1))];

        -- Gera preço aleatório entre 20 e 150
        preco := 20 + (random() * 130);

        -- Gera quantidade aleatória entre 1 e 100
        quantidade := 1 + floor(random() * 100);

        -- Seleciona uma loja aleatória
        loja := lojas[1 + floor(random() * array_length(lojas, 1))];

        -- Insere o produto
        INSERT INTO products (store_id, name, description, price, ammount)
        VALUES (loja, nome, descricao, ROUND(preco::numeric, 2), quantidade);

        -- Feedback a cada 10000 inserções
        IF i % 10000 = 0 THEN
            RAISE NOTICE 'Inseridos % produtos', i;
        END IF;
    END LOOP;
END $$;

-- Adiciona um índice para melhorar algumas buscas (opcional)
CREATE INDEX IF NOT EXISTS idx_products_name ON products(name);
CREATE INDEX IF NOT EXISTS idx_products_description ON products(description);

-- Analisa a tabela para atualizar as estatísticas do planejador de consultas
ANALYZE products;
