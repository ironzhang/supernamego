if MatchFuncs == nil then
	MatchFuncs = {}
end

function PrintTable(n, t)
	print(n, ":")
	for k, v in pairs(t) do
		print(k, v)
	end
end

MatchFuncs["www.test2.com"] = function(ctx, clusters)
	UserDefineLablesMatch = function(ctx, labels)
		if ctx["X-Zone"] ~= labels["X-Zone"] then
			return false
		end
		if ctx["X-Lane"] ~= labels["X-Lane"] then
			return false
		end
		return true
	end

	DefaultLabelsMatch = function(ctx, labels)
		if ctx["X-Zone"] ~= labels["X-Zone"] then
			return false
		end
		if labels["X-Lane"] ~= "default" then
			return false
		end
		return true
	end

	for _, cluster in pairs(clusters) do
		if UserDefineLablesMatch(ctx, cluster.Labels) then
			return { {cluster.Name, 1} }
		end
	end

	for _, cluster in pairs(clusters) do
		if DefaultLabelsMatch(ctx, cluster.Labels) then
			return { {cluster.Name, 1} }
		end
	end

	return {}
end
